package internal

import (
	"fmt"
	"strings"

	"github.com/Danissimode/Palto/config"
	"github.com/Danissimode/Palto/system"
)

func (m *Manager) GetPaltoPanes() ([]system.PaltoPaneDetails, error) {
	currentPaneId, _ := system.PaltoCurrentPaneId()
	windowTarget, _ := system.PaltoCurrentWindowTarget()
	currentPanes, _ := system.PaltoPanesDetails(windowTarget)

	for i := range currentPanes {
		currentPanes[i].IsPaltopalsPane = currentPanes[i].Id == currentPaneId
		currentPanes[i].IsPaltopalsExecPane = currentPanes[i].Id == m.ExecPane.Id
		currentPanes[i].IsPrepared = currentPanes[i].Id == m.ExecPane.Id
		if currentPanes[i].IsSubShell {
			currentPanes[i].OS = "OS Unknown (subshell)"
		} else {
			currentPanes[i].OS = m.OS
		}

	}
	return currentPanes, nil
}

func (m *Manager) GetPaltoPanesInXml(config *config.Config) string {
	currentPaltoWindow := strings.Builder{}
	currentPaltoWindow.WriteString("<current_Palto_window_state>\n")
	panes, _ := m.GetPaltoPanes()

	// Filter out Paltopals_pane
	var filteredPanes []system.PaltoPaneDetails
	for _, p := range panes {
		if !p.IsPaltopalsPane {
			filteredPanes = append(filteredPanes, p)
		}
	}
	for _, pane := range filteredPanes {
		if !pane.IsPaltopalsPane {
			pane.Refresh(m.GetMaxCaptureLines())
		}
		if pane.IsPaltopalsExecPane {
			m.ExecPane = &pane
		}

		var title string
		if pane.IsPaltopalsExecPane {
			title = "Paltopals_exec_pane"
		} else {
			title = "read_only_pane"
		}

		currentPaltoWindow.WriteString(fmt.Sprintf("<%s>\n", title))
		currentPaltoWindow.WriteString(fmt.Sprintf(" - Id: %s\n", pane.Id))
		currentPaltoWindow.WriteString(fmt.Sprintf(" - CurrentPid: %d\n", pane.CurrentPid))
		currentPaltoWindow.WriteString(fmt.Sprintf(" - CurrentCommand: %s\n", pane.CurrentCommand))
		currentPaltoWindow.WriteString(fmt.Sprintf(" - CurrentCommandArgs: %s\n", pane.CurrentCommandArgs))
		currentPaltoWindow.WriteString(fmt.Sprintf(" - Shell: %s\n", pane.Shell))
		currentPaltoWindow.WriteString(fmt.Sprintf(" - OS: %s\n", pane.OS))
		currentPaltoWindow.WriteString(fmt.Sprintf(" - LastLine: %s\n", pane.LastLine))
		currentPaltoWindow.WriteString(fmt.Sprintf(" - IsActive: %d\n", pane.IsActive))
		currentPaltoWindow.WriteString(fmt.Sprintf(" - IsPaltopalsPane: %t\n", pane.IsPaltopalsPane))
		currentPaltoWindow.WriteString(fmt.Sprintf(" - IsPaltopalsExecPane: %t\n", pane.IsPaltopalsExecPane))
		currentPaltoWindow.WriteString(fmt.Sprintf(" - IsPrepared: %t\n", pane.IsPrepared))
		currentPaltoWindow.WriteString(fmt.Sprintf(" - IsSubShell: %t\n", pane.IsSubShell))
		currentPaltoWindow.WriteString(fmt.Sprintf(" - HistorySize: %d\n", pane.HistorySize))
		currentPaltoWindow.WriteString(fmt.Sprintf(" - HistoryLimit: %d\n", pane.HistoryLimit))

		if !pane.IsPaltopalsPane && pane.Content != "" {
			currentPaltoWindow.WriteString("<pane_content>\n")
			currentPaltoWindow.WriteString(pane.Content)
			currentPaltoWindow.WriteString("\n</pane_content>\n")
		}

		currentPaltoWindow.WriteString(fmt.Sprintf("</%s>\n\n", title))
	}

	currentPaltoWindow.WriteString("</current_Palto_window_state>\n")
	return currentPaltoWindow.String()
}
