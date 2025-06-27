#!/bin/bash

set -e

PROJECT_NAME=$1

if [ -z "$PROJECT_NAME" ]; then
  echo "Usage: ./create-vite-custom.sh <project-name>"
  exit 1
fi

export NVM_DIR="$HOME/.nvm"
if [ -s "$NVM_DIR/nvm.sh" ]; then
  . "$NVM_DIR/nvm.sh"
else
  echo "❌ nvm not found. Please install nvm first."
  exit 1
fi

CURRENT_NODE_VERSION=$(node -v)

if [[ "$CURRENT_NODE_VERSION" != "v22"* ]]; then
  echo "⚠️ Current Node.js version is not 22. Switching to Node.js 22 using nvm..."
  nvm use 22
else
  echo "✅ Node.js version 22 is already in use."
fi

npm create vite@latest "$PROJECT_NAME" -- --template react

cd "$PROJECT_NAME"

rm -rf src/App.css src/index.css src/App.jsx 
rm -f public/vite.svg
rm -rf src/assets

cat > src/App.jsx <<EOF
export default function App() {
  return <h1>App Component!</h1>;
}
EOF

cat > src/main.jsx <<EOF
import { StrictMode } from 'react';
import ReactDOM from 'react-dom/client';
import App from './App.jsx';

ReactDOM.createRoot(document.getElementById('root')).render(
  <StrictMode>
    <App />
  </StrictMode>
);
EOF

cat > index.html <<EOF
<!doctype html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <link rel="stylesheet" href="/src/index.css" />
  </head>
  <body>
    <div id="root"></div>
    <script type="module" src="/src/main.jsx"></script>
  </body>
</html>
EOF

cat > src/index.css <<EOF
* {
  box-sizing: border-box;
}

body {
  margin: 0;
}
EOF

cat > vite.config.js <<EOF
import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
  host: true,
  watch: {
    server: {
        usePolling: true
    }
  }
})
EOF

# git init
# git add .
# git commit -m "Initial commit from custom Vite script"

echo "✅ Project '$PROJECT_NAME' created!"