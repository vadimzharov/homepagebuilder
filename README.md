## Home page builder 

Tool to create home page for Home lab based on HTML code generated by ChatGPT.

To use:

```
git clone https://github.com/vadimzharov/homepagebuilder.git
cd homepagebuilder
copy config-example.yaml config.yaml
```

Update config.yaml (list of apps, provide ChatGPT API key). Then:

```
go run .
```

Home page will be available `http://localhost:8080`

To regenerate artifacts run

`curl localhost:8080/api/regenerateall` - to regenerate all 

`curl localhost:8080/api/regeneratemain` - to regenerate main page layout

`curl localhost:8080/api/regeneratepanels` - to regenerate panels layout

`curl localhost:8080/api/regenerateapps` - to regenerate apps description

