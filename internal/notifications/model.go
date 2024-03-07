package notifications

var notificationsSlice = []Notification{
	{
		Title:       "Whispers of the Celestial Symphony: A Sci-Fi Epic",
		Description: "Embark on a thrilling journey through the cosmos as a disparate group of explorers uncovers the secrets behind the enigmatic Celestial Symphony, a cosmic force that could reshape the fate of galaxies.",
	},
	{
		Title:       "Chronicles of the Enchanted Quill: Fantasy Adventures Unveiled",
		Description: "Dive into a world of magic and wonder, where an ancient quill holds the key to unlocking portals to fantastical realms, and a group of unlikely heroes must navigate through perilous landscapes to save their world from impending doom.",
	},
	{
		Title:       "Beneath Neon Skies: A Cyberpunk Mystery",
		Description: "In the neon-drenched streets of a futuristic metropolis, a lone detective must unravel a web of corruption, deceit, and high-tech intrigue, all while confronting the ghosts of a city that never sleeps.",
	},
	{
		Title:       "Harvest Moon Serenade: A Romantic Tale of Rural Bliss",
		Description: "Experience the charm of a picturesque countryside as two souls find love amidst rolling fields and quaint farm life, exploring the profound beauty of nature and the enduring power of connection.",
	},
	{
		Title:       "Quantum Convergence: Exploring the Frontiers of Science and Reality",
		Description: "Join a team of brilliant scientists on a groundbreaking expedition into the uncharted territories of quantum physics, where reality bends, time warps, and the very fabric of existence is questioned in a quest for knowledge that transcends the boundaries of conventional understanding.",
	},
}

type Notification struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
