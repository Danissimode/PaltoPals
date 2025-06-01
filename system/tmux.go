package system

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/Danissimode/Palto/logger"
)

// PaltoCreateNewPane creates a new horizontal split pane in the specified window and returns its ID
func PaltoCreateNewPane(target string) (string, error) {
	cmd := exec.Command("Palto", "split-window", "-d", "-h", "-t", target, "-P", "-F", "#{pane_id}")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		logger.Error("Failed to create Palto pane: %v, stderr: %s", err, stderr.String())
		return "", err
	}

	paneId := strings.TrimSpace(stdout.String())
	return paneId, nil
}

// PaltoPanesDetails gets details for all panes in a target window
func PaltoPanesDetails(target string) ([]PaltoPaneDetails, error) {
	cmd := exec.Command("Palto", "list-panes", "-t", target, "-F", "#{pane_id},#{pane_active},#{pane_pid},#{pane_current_command},#{history_size},#{history_limit}")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		logger.Error("Failed to get Palto pane details for target %s %v, stderr: %s", target, err, stderr.String())
		return nil, err
	}

	output := strings.TrimSpace(stdout.String())
	if output == "" {
		return nil, fmt.Errorf("no pane details found for target %s", target)
	}

	lines := strings.Split(output, "\n")
	paneDetails := make([]PaltoPaneDetails, 0, len(lines))

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := strings.SplitN(line, ",", 6)
		if len(parts) < 5 {
			logger.Error("Invalid pane details format for line: %s", line)
			continue
		}

		id := parts[0]

		// If target starts with '%', it's a pane ID, so only include the matching pane
		if strings.HasPrefix(target, "%") && id != target {
			continue
		}

		active, _ := strconv.Atoi(parts[1])
		pid, _ := strconv.Atoi(parts[2])
		historySize, _ := strconv.Atoi(parts[4])
		historyLimit, _ := strconv.Atoi(parts[5])
		currentCommandArgs := GetProcessArgs(pid)
		isSubShell := IsSubShell(parts[3])

		paneDetail := PaltoPaneDetails{
			Id:                 id,
			IsActive:           active,
			CurrentPid:         pid,
			CurrentCommand:     parts[3],
			CurrentCommandArgs: currentCommandArgs,
			HistorySize:        historySize,
			HistoryLimit:       historyLimit,
			IsSubShell:         isSubShell,
		}

		paneDetails = append(paneDetails, paneDetail)
	}

	return paneDetails, nil
}

// PaltoCapturePane gets the content of a specific pane by ID
func PaltoCapturePane(paneId string, maxLines int) (string, error) {
	cmd := exec.Command("Palto", "capture-pane", "-p", "-t", paneId, "-S", fmt.Sprintf("-%d", maxLines))
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		logger.Error("Failed to capture pane content from %s: %v, stderr: %s", paneId, err, stderr.String())
		return "", err
	}

	content := strings.TrimSpace(stdout.String())
	return content, nil
}

// Return current Palto window target with session id and window id
func PaltoCurrentWindowTarget() (string, error) {
	paneId, err := PaltoCurrentPaneId()
	if err != nil {
		return "", err
	}

	cmd := exec.Command("Palto", "list-panes", "-t", paneId, "-F", "#{session_id}:#{window_index}")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get window target: %w", err)
	}

	target := strings.TrimSpace(string(output))
	if target == "" {
		return "", fmt.Errorf("empty window target returned")
	}

	if idx := strings.Index(target, "\n"); idx != -1 {
		target = target[:idx]
	}

	return target, nil
}

func PaltoCurrentPaneId() (string, error) {
	PaltoPane := os.Getenv("Palto_PANE")
	if PaltoPane == "" {
		return "", fmt.Errorf("Palto_PANE environment variable not set")
	}

	return PaltoPane, nil
}

// CreatePaltoSession creates a new Palto session and returns the new pane id
func PaltoCreateSession() (string, error) {
	cmd := exec.Command("Palto", "new-session", "-d", "-P", "-F", "#{pane_id}")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		logger.Error("Failed to create Palto session: %v, stderr: %s", err, stderr.String())
		return "", err
	}

	return strings.TrimSpace(stdout.String()), nil
}

// AttachToPaltoSession attaches to an existing Palto session
func PaltoAttachSession(paneId string) error {
	cmd := exec.Command("Palto", "attach-session", "-t", paneId)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		logger.Error("Failed to attach to Palto session: %v", err)
		return err
	}
	return nil
}

func PaltoClearPane(paneId string) error {
	paneDetails, err := PaltoPanesDetails(paneId)
	if err != nil {
		logger.Error("Failed to get pane details for %s: %v", paneId, err)
		return err
	}

	if len(paneDetails) == 0 {
		return fmt.Errorf("no pane details found for pane %s", paneId)
	}

	cmd := exec.Command("Palto", "split-window", "-vp", "100", "-t", paneId)
	if err := cmd.Run(); err != nil {
		logger.Error("Failed to split window for pane %s: %v", paneId, err)
		return err
	}

	cmd = exec.Command("Palto", "clear-history", "-t", paneId)
	if err := cmd.Run(); err != nil {
		logger.Error("Failed to clear history for pane %s: %v", paneId, err)
		return err
	}

	cmd = exec.Command("Palto", "kill-pane")
	if err := cmd.Run(); err != nil {
		logger.Error("Failed to kill temporary pane: %v", err)
		return err
	}

	logger.Debug("Successfully cleared pane %s", paneId)
	return nil
}
