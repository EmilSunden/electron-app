const { app, BrowserWindow, shell} = require('electron');
const path = require('path');

async function createWindow() {
  const isDev = (await import('electron-is-dev')).default;

  const win = new BrowserWindow({
    width: 800,
    height: 600,
    webPreferences: {
      preload: path.join(__dirname, 'preload.js'),
      contextIsolation: true, // Ensure this is set to false if you need Node.js integration in your renderer process
    },
  });

  win.loadURL(
    isDev
      ? 'http://localhost:3000'
      : `file://${path.join(__dirname, '../build/index.html')}`
  );

  win.webContents.on('will-navigate', (e, url) => {
    const currentUrl = win.webContents.getURL();
    const currentOrigin = new URL(currentUrl).origin;
    const targetOrigin = new URL(url).origin;

    if(currentOrigin !== targetOrigin) {
      e.preventDefault();
      shell.openExternal(url);
    }
  });

  win.webContents.setWindowOpenHandler(({url}) => {
    shell.openExternal(url);
    return {action : 'deny'}
  })
}

app.whenReady().then(createWindow)

app.on('window-all-closed', () => {
  if (process.platform !== 'darwin') {
    app.quit();
  }
});
