package pagebuilder

const mainpagequery string = `You are a web developer. As a web developer, you need to generate HTML code to a create a web page, and the main purpose of this web page is to provide list of applications available and running in my Kubernetes cluster. 
The page will have a list of panels, where one panel represents one application running in my Kubernetes cluster.
As the first answer, provide me HTML code for the main page, but do not include HTML code for each panel. 
Main page should follow the requirements:	
1. Title 'Applications running in my K8S cluster'	
2. Have a fancy theme, include fancy theme CSS styles in the generated HTML code
3. There can be multiple panels in one line.
4. Insert comment showing where I can paste panels code, comment should have phrase "INSERT PANELS CODE HERE"
5. Have a grid layout.
6. Have a button on the bottom calling api endpoint "/api/regenerateall". The button name is "Regenerate all pages"`

// const panelquery string = `You are a web developer. As a web developer, you need to generate HTML code for a item describing an application. This item's code will be inserted to the mane page with grid layout.
// Requirements for the item:
// * The item title must be text "APPLICATION_NAME"
// * The item must have a detailed description of the application. Use phrase "DETAILED_DESCRIPTION" as text for this field.
// * The item should provide URL where this application is hosted. Use URL "http://url_for_the_application" as a URL.
// * The item should have shades
// * As an image for the item use image from URL "http://url_for_the_image" and size of the image must be set as width = 150 and height = 150
// `

const panelquery string = `You are a web developer. As a web developer, you need to generate Golang HTML Template for an item describing an application. This item's code will be inserted to the main page with the grid layout.
Requirements for the item/generated Golang HTML template:
* The item title set as {{.Name}}
* The item must have a detailed description of the application. Use {{.Description}} in the template for this.
* The item should provide URL where this application is hosted. Use {{.URL}} in the template for this
* The item should have shades
* Image for the item must be optional if {{.Image}} set to true. If image is defined, use image from URL set as {{.ImageURL}} in the template and size of the image must be set as width = {{.ImageWidth}} and height = {{.ImageHeight}}`

const appsreqs string = `For each application in the list provide description. The output must be in JSON format - [{"name": <application name>, "description": <application description>}]. As a response provide only one JSON containing everything.
Applications:
`

const apptemplate string = "Application - {{.Name}} \n"
