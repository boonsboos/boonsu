package commands

import "github.com/bwmarrin/discordgo"

type Command = func(*discordgo.Session, *discordgo.Message, []string)

/*
example command file

```
package commands

import "github.com/bwmarrin/discordgo"

func nameOfCommand(s *discordgo.Session, m *discordgo.Message, c []string)
```
*/

// oh yippee-kay-yay
var all_commands = map[string]Command{
	"ping":    pingCommand,
	"linkosu": osuLinkCommand,
	"stats":   statsCommand,
	"osu":     osuProfileCommand,
	"osutop":  osuTopCommand,
	"rs":      osuRecentCommand,
}
