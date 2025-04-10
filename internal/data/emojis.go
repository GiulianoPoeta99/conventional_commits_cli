// Package data provides functions to retrieve commit types and emojis.
// GetEmojis returns a slice of Emoji structures representing the available emojis
// for enhancing commit messages with visual indicators.
package data

import (
	t "github.com/GiulianoPoeta99/conventional_commits_cli/internal/types"
)

// GetEmojis returns a slice of Emoji.
// Each Emoji contains a symbol, a code, and a brief description of its intended usage.
func GetEmojis() []t.Emoji {
	return []t.Emoji{
		{
			Symbol:      "🎨",
			Code:        "art",
			Description: "Improve structure / format of the code",
		},
		{
			Symbol:      "⚡",
			Code:        "zap",
			Description: "Improve performance",
		},
		{
			Symbol:      "🔥",
			Code:        "fire",
			Description: "Remove code or files",
		},
		{
			Symbol:      "🐛",
			Code:        "bug",
			Description: "Fix a bug",
		},
		{
			Symbol:      "🚑",
			Code:        "ambulance",
			Description: "Critical hotfix",
		},
		{
			Symbol:      "✨",
			Code:        "sparkles",
			Description: "Introduce new features",
		},
		{
			Symbol:      "📝",
			Code:        "memo",
			Description: "Add or update documentation",
		},
		{
			Symbol:      "🚀",
			Code:        "rocket",
			Description: "Deploy stuff",
		},
		{
			Symbol:      "💄",
			Code:        "lipstick",
			Description: `Add or update the UI and style files`,
		},
		{
			Symbol:      "🎉",
			Code:        "tada",
			Description: "Begin a project",
		},
		{
			Symbol:      "✅",
			Code:        "white_check_mark",
			Description: "Add, update, or pass test",
		},
		{
			Symbol:      "🔓",
			Code:        "lock",
			Description: "Fix security issues",
		},
		{
			Symbol:      "🔐",
			Code:        "closed_lock_with_key",
			Description: "Add or update secrets",
		},
		{
			Symbol:      "🔖",
			Code:        "bookmark",
			Description: "Release / version tags",
		},
		{
			Symbol:      "🚨",
			Code:        "rotating_light",
			Description: "Fix compiler / linter warnings",
		},
		{
			Symbol:      "🚧",
			Code:        "construction",
			Description: "Work in progress",
		},
		{
			Symbol:      "💚",
			Code:        "green_heart",
			Description: "Fix CI build",
		},
		{
			Symbol:      "⬇️",
			Code:        "arrow_down",
			Description: "Downgrade dependencies",
		},
		{
			Symbol:      "⬆️",
			Code:        "arrow_up",
			Description: "Upgrade dependencies",
		},
		{
			Symbol:      "📌",
			Code:        "pushpin",
			Description: "Pin dependencies to specific versions",
		},
		{
			Symbol:      "👷",
			Code:        "construction_worker",
			Description: "Add or update CI build system",
		},
		{
			Symbol:      "📈",
			Code:        "chart_with_upwards_trend",
			Description: "Add or update analytics or track code",
		},
		{
			Symbol:      "♻️",
			Code:        "recycle",
			Description: "Refactor code",
		},
		{
			Symbol:      "➕",
			Code:        "heavy_plus_sign",
			Description: "Add a dependency",
		},
		{
			Symbol:      "➖",
			Code:        "heavy_minus_sign",
			Description: "Remove a dependency",
		},
		{
			Symbol:      "🔧",
			Code:        "wrench",
			Description: "Add or update configuration files",
		},
		{
			Symbol:      "🔨",
			Code:        "hammer",
			Description: "Add or update development scripts",
		},
		{
			Symbol:      "🌐",
			Code:        "globe_with_meridians",
			Description: "Internationalization and localization",
		},
		{
			Symbol:      "✏️",
			Code:        "pencil2",
			Description: "Fix typos",
		},
		{
			Symbol:      "💩",
			Code:        "poop",
			Description: "Write bad code that needs to be improved",
		},
		{
			Symbol:      "⏪",
			Code:        "rewind",
			Description: "Revert changes",
		},
		{
			Symbol:      "🔀",
			Code:        "twisted_rightwards_arrows",
			Description: "Merge branches",
		},
		{
			Symbol:      "📦",
			Code:        "package",
			Description: "Add or update compiled files or packages",
		},
		{
			Symbol:      "👽",
			Code:        "alien",
			Description: "Update code due to external API changes",
		},
		{
			Symbol:      "🚚",
			Code:        "truck",
			Description: "Move or rename resources (e.g.: files, paths, routes)",
		},
		{
			Symbol:      "📄",
			Code:        "page_facing_up",
			Description: "Add or update license",
		},
		{
			Symbol:      "💥",
			Code:        "boom",
			Description: "Introduce breaking changes",
		},
		{
			Symbol:      "🍱",
			Code:        "bento",
			Description: "Add or update assets",
		},
		{
			Symbol:      "♿",
			Code:        "wheelchair",
			Description: "Improve accessibility",
		},
		{
			Symbol:      "💡",
			Code:        "bulb",
			Description: "Add or update comments in source code",
		},
		{
			Symbol:      "🍻",
			Code:        "beers",
			Description: "Write code drunkenly",
		},
		{
			Symbol:      "💬",
			Code:        "speech_balloon",
			Description: "Add or update text and literals",
		},
		{
			Symbol:      "🗃️",
			Code:        "card_file_box",
			Description: "Perform database related changes",
		},
		{
			Symbol:      "🔊",
			Code:        "loud_sound",
			Description: "Add or update logs",
		},
		{
			Symbol:      "🔇",
			Code:        "mute",
			Description: "Remove logs",
		},
		{
			Symbol:      "👥",
			Code:        "busts_in_silhouette",
			Description: "Add or update contributor(s)",
		},
		{
			Symbol:      "🚸",
			Code:        "children_crossing",
			Description: "Improve user experience / usability",
		},
		{
			Symbol:      "🏗️",
			Code:        "building_construction",
			Description: "Make architectural changes",
		},
		{
			Symbol:      "📱",
			Code:        "iphone",
			Description: "Work on responsive design",
		},
		{
			Symbol:      "🤡",
			Code:        "clown_face",
			Description: "Mock things",
		},
		{
			Symbol:      "🥚",
			Code:        "egg",
			Description: "Add or update an easter egg",
		},
		{
			Symbol:      "🙈",
			Code:        "see_no_evil",
			Description: "Add or update a .gitignore file",
		},
		{
			Symbol:      "📸",
			Code:        "camera_flash",
			Description: "Add or update snapshots",
		},
		{
			Symbol:      "⚗️",
			Code:        "alembic",
			Description: "Perform experiments",
		},
		{
			Symbol:      "🔍",
			Code:        "mag",
			Description: "Improve SEO",
		},
		{
			Symbol:      "🏷️",
			Code:        "label",
			Description: "Add or update types",
		},
		{
			Symbol:      "🌱",
			Code:        "seedling",
			Description: "Add or update seed files",
		},
		{
			Symbol:      "🚩",
			Code:        "triangular_flag_on_post",
			Description: "Add, update, or remove feature flags",
		},
		{
			Symbol:      "🥅",
			Code:        "goal_net",
			Description: "Catch errors",
		},
		{
			Symbol:      "💫",
			Code:        "dizzy",
			Description: "Add or update animations and transitions",
		},
		{
			Symbol:      "🗑️",
			Code:        "wastebasket",
			Description: "Deprecate code that needs to be cleaned up",
		},
		{
			Symbol:      "🛂",
			Code:        "passport_control",
			Description: "Work on code related to authorization, roles, and permissions",
		},
		{
			Symbol:      "🩹",
			Code:        "adhesive_bandage",
			Description: "Simple fix for a non-critical issue",
		},
		{
			Symbol:      "🧐",
			Code:        "monocle_face",
			Description: "Data exploration / inspection",
		},
		{
			Symbol:      "⚰️",
			Code:        "coffin",
			Description: "Remove dead code",
		},
		{
			Symbol:      "🧪",
			Code:        "test_tube",
			Description: "Add a failing test",
		},
		{
			Symbol:      "👔",
			Code:        "necktie",
			Description: "Add or update business logic",
		},
		{
			Symbol:      "🩺",
			Code:        "stethoscope",
			Description: "Add or update health check",
		},
		{
			Symbol:      "🧱",
			Code:        "bricks",
			Description: "Infrastructure related changes",
		},
		{
			Symbol:      "🧑‍💻",
			Code:        "technologist",
			Description: "Improve developer experience",
		},
		{
			Symbol:      "💸",
			Code:        "money_with_wings",
			Description: "Add sponsorships or money related infrastructure",
		},
		{
			Symbol:      "🧵",
			Code:        "thread",
			Description: "Add or update code related to multithreading or concurrency",
		},
		{
			Symbol:      "🦺",
			Code:        "safety_vest",
			Description: "Add or update code related to validation",
		},
	}
}
