import { Route, Routes } from "react-router";
import "./App.css";
import Home from "./pages/Home";
import OAuthLoginCallback from "./pages/OAuthLoginCallback";

const App = () => {
  return (
    <Routes>
      <Route path="/" element={<Home />} />
      <Route path="/callback" element={<OAuthLoginCallback />} />
    </Routes>
  );
};

export default App;
