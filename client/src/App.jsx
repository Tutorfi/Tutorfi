
function Home () {
  return (
    <div>
      <h1>Home testing</h1>
      <h2>Home testing</h2>
    </div>
  )
}

function App () {
  return (
    <div>
      <h1>App testing</h1>
      <h2>App testing</h2>
      <Home />
    </div>
  )
}

export default App
