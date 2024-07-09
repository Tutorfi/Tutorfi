### client
This is the client side app in the Tutorfi container for frontend using vite and solidjs.

### For frontend devs
*** The commands here run under the client folder.

`npm install` to get node_module
`npm run dev` or `npm start` to test
Open [http://localhost:3000](http://localhost:3000) to view it in the browser.

`npx eslint "src/**/*.jsx"` to lint the jsx files locally and manually

NOTE:
We don't need to `npm run build` for development so make sure when you docker compose, the dist folder is not in your directory.
Also we don't need the node_modules folder but it can be there when docker compose since it is ignored in docker compose.
We are not using vite-plugin-eslint for development because this causes issues in deployment.

### Deployment
Docker compose the entire container to deploy, also need to do this to test frontend-to-backend interactions.
cd to root dir `Tutorfi/`
Run: 
docker compose -f ./deployments/docker-compose.yml -f ./deployments/docker-compose.dev.yml up -d