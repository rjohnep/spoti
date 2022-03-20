import useLocalStorage from "@rehooks/local-storage";
import React from "react";
import Dashboard from "../Dashboard";
import logo from "./logo.svg";

const Home = () => {
  const [accessToken] = useLocalStorage("accessToken", undefined);

  return (
    <>
      {!accessToken && (
        <div className="App">
          <header className="App-header">
            <a href="/login">Login</a>
            <img src={logo} className="App-logo" alt="logo" />
          </header>
        </div>
      )}
      {accessToken && <Dashboard />}
    </>
  );
};

export default Home;
