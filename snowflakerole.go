package main

import (
	"github.com/bwmarrin/discordgo"
	"strconv"
	"strings"
)

func SnowflakeRole(s *discordgo.Session, guildID, userID, color, name string) {

	var role *discordgo.Role
	memberData, err := s.GuildMember(guildID, userID)
	memberCurrentRoles := memberData.Roles

	if len(memberCurrentRoles) > 0 && memberCurrentRoles[0] != "gay baby jailed" {
		role, err = s.GuildRoleEdit(guildID, memberCurrentRoles[0], name, ParseColorInput(color), false, 0, false)
		HandleError(err, "Error making changes to role with ID "+memberCurrentRoles[0])
		return
	}

	role, err = s.GuildRoleCreate(guildID)
	HandleError(err, "Error creating role")

	role, err = s.GuildRoleEdit(guildID, role.ID, name, ParseColorInput(color), role.Hoist, role.Permissions, false)
	HandleError(err, "Error configuring role with ID "+role.ID)

	err = s.GuildMemberRoleAdd(guildID, userID, role.ID)
	HandleError(err, "Error adding role to user with ID "+userID)

	return
}

func ParseColorInput(color string) (parsed int) {
	color = strings.Replace(color, "#", "", -1)
	colorInt, _ := strconv.ParseInt(color, 16, 64)
	parsed = int(colorInt)

	return parsed
}
