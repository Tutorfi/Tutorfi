import { Routes, Route } from "@solidjs/router";
import Navbar from "./components/Navbar";
function Home() {
  return (
    <div>
      <h1>Home testing</h1>
      <h2>Home testing</h2>
    </div>
  );
}

function App() {
  return (
    <>
      <Routes>
        <Route path="/" element={<Home />} />
      </Routes>
    </>
  );
}

export default App;
