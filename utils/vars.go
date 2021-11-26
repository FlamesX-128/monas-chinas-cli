package utils

var (
	BaseURL string = "https://monoschinos2.com/"
)

var HTML = [5]string{`
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>
`, `</title>
    <script src="https://monoschinos2.com/public/js/jquery.min.js"></script>
		<style>
			:root {
				--background-color: #010409;
			}
			
			* {
				color: #ccc;
			}
			
			html, body {
				background: var(--background-color);
				background-color: var(--background-color);
			
				height: 100vh;
				overflow: hidden;
				width: 100%;
			}
			
			div#appPopUp {
				display: grid;
				grid-template-columns: 100%;
				grid-template-rows: auto;
				grid-template-areas: "mainContainer";
			
				align-content: center;
				background: var(--background-color);
				background-color: var(--background-color);
				height: 100vh;
				justify-content: center;
				position: fixed;
				width: 100%;
			}
			
			div#appPopUp div.mainContainer {
				display: grid;
				grid-area: mainContainer;
				grid-template-columns: 100%;
				grid-template-rows: auto;
				grid-template-areas: "loading" "desc";
			
				align-items: center;
				background: var(--background-color);
				background-color: var(--background-color);
				justify-content: center;
			}
			
			
			div#appPopUp div.mainContainer div.loading {
				align-self: center;
				justify-self: center;
				-webkit-animation: spin 1s linear infinite; 
				animation: spin 1s linear infinite;
			
				border: 16px solid #f3f3f3;
				border-radius: 50%;
				border-top: 16px solid #3498db;
				grid-area: loading;
				width: 120px;
				height: 120px;
			}
			
			div#appPopUp div.mainContainer p {
				grid-area: desc;
				text-align: center;
				margin-top: 30px;
			}
			
			@-webkit-keyframes spin {
				0% { -webkit-transform: rotate(0deg); }
				100% { -webkit-transform: rotate(360deg); }
			}
			
			@keyframes spin {
				0% { transform: rotate(0deg); }
				100% { transform: rotate(360deg); }
			}
		</style>
  </head>

  <body>
		<div id="appPopUp">
			<div class="mainContainer">
				<div class="loading"></div>
				<p>Obteniendo episodio</p>
			</div>
		</div>

    <div class="heromain">
      <div class="playother">
`, `	</div>

			<div class="heroarea">
        <div class="playmain">
          <div class="player"></div>

          <div class="playdiv">
            <div class="playcontrol">
              <div class="dropdown playdwn">
                <ul class="dropdown-menu dropcap">
`, `						</ul>
              </div>
            </div>
          </div>
        </div>
      </div>

    </div>

    <script src="https://monoschinos2.com/public/js/capitulo.js"></script>

		<script type="module">
			function getPage(url) {
				const animeURL = url.split("?url=");
				
				window.location.href = (animeURL[1] !== undefined)
					? animeURL[1]
					: animeURL[0];
			}

			await (function main(toReturn) {
				setTimeout(function () {
					if (toReturn && toReturn.length) getPage(toReturn);

					const services = document.getElementsByClassName("play-video");

					for (let i = 0; i < services.length; i++)
						if (services[i].innerText === "`, `")
							document.getElementsByClassName("play-video")[i].click();


					setTimeout(function () {
						const url = document.querySelector('iframe.embed-responsive-item');

						if (!url) return main();

						main(url.src);
					}, 500);
				}, 2500)
			})();

			export {};
		</script>
  </body>
</html>
`}
