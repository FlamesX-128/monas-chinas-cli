const trayMenuItems = {
	VERSION: () => Neutralino.os.showMessageBox(
		"Version information",
		`Neutralinojs server: v${NL_VERSION} | Neutralinojs client: v${NL_CVERSION}`
	),
	QUIT: () => Neutralino.app.exit()
};

(async function main() {
	if (NL_OS != "Darwin" && NL_MODE == "window") Neutralino.os.setTray({
		icon: "/resources/icons/trayIcon.png",
		menuItems: [
			{ id: "VERSION", text: "Get version" },
			{ id: "SEP", text: "-" },
			{ id: "QUIT", text: "Quit" }
		]
	});


	Neutralino.events.on("windowClose", () => Neutralino.app.exit());
	Neutralino.events.on("trayMenuItemClicked", (event) => trayMenuItems(event.detail.id));


	Neutralino.init();

	const episodeContainer = document.getElementById("episodeContainer");

	const res = await (await fetch("http://localhost:43078/")).json()
		.catch((err) => episodeContainer.innerHTML = `${err}`);

	episodeContainer.innerHTML = `
		<iframe src="${res.Url}" height="100%" width="100%"></iframe>
	`;
})();