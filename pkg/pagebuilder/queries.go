package pagebuilder

const Defaultmainpagequery string = `You are a web developer. As a web developer, you need to generate HTML code to a create a web page, and the main purpose of this web page is to provide list of applications available and running in my Kubernetes cluster. 
The page will have a list of panels, where one panel represents one application running in my Kubernetes cluster.
As the first answer, provide me HTML code for the main page, but do not include HTML code for each panel. 
Main page should follow the requirements:	
1. Title 'Applications running in my K8S cluster'	
2. Have a fancy theme, include fancy theme CSS styles in the generated HTML code
3. There can be multiple panels in one line.
4. Insert comment showing where I can paste panels code, comment should have phrase "INSERT PANELS CODE HERE"
5. Have a grid layout.
6. Have buttons on the bottom, each of them is calling Javascript to call API endpoint. Include Javascript code in the generated HTML code. List of the buttons:
* Button title - "Regenerate all assets" and it should call API endpoint "/api/regenerateall" via Javascript code included in the generated HTML code
* Button title - "Regenerate main page layout" and it should call API endpoint "/api/regeneratemain" via Javascript code included in the generated HTML code
* Button title - "Regenerate panels layout" and it should call API endpoint "/api/regeneratepanels" via Javascript code included in the generated HTML code
* Button title - "Regenerate apps description" and it should call API endpoint "/api/regenerateapps" via Javascript code included in the generated HTML code
`

// 6. Have a button on the bottom calling Javascript to call API endpoint "/api/regenerateall". Include Javascript code in the generated HTML code. The button name is "Regenerate all pages"
const Defaultpanelquery string = `You are a web developer. As a web developer, you need to generate Golang HTML Template for an item describing an application. This item's code will be inserted to the main page with the grid layout.
Requirements for the item/generated Golang HTML template:
* The item title set as {{.Name}}
* The item must have fancy theme. Include fancy theme CSS styles in the generated HTML template.
* The item must have a detailed description of the application. Use {{.Description}} in the template for this.
* The item should provide URL where this application is hosted. Use {{.URL}} in the template for this. The URL should be activated by pressing the button and opens in a new tab.
* The item should be framed and have shades.
* Image for the item must be optional if {{.Image}} set to true. If image is defined, use image from URL set as {{.ImageURL}} in the template and size of the image must be set as width = {{.ImageWidth}} and height = {{.ImageHeight}}
Do not use any other variables in the Golang HTML template except listed above.
`

const appsreqs string = `For each application in the list provide description. The output must be in JSON format - [{"name": <application name>, "description": <application description>}]. As a response provide only one JSON containing everything.
Applications:
`

const apptemplate string = "Application - {{.Name}} \n"
