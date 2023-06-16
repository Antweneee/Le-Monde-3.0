import React, { useEffect }  from 'react';
import ReactDOM from 'react-dom';
import { CookiesProvider } from "react-cookie";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import './index.css';
import NoPage from "./pages/NoPage";
import Layout from './pages/Layout';
import Home from "./pages/Home";
import Login from "./pages/Login";
import AccountCreation from "./pages/AccountCreation";

export default function App() {
  useEffect(() => {
  }, [])
  
  return (
    <BrowserRouter>
      <Routes>
      <Route index path="/Login" element={<Login />} />
      <Route path="/Register" element={<AccountCreation />} />
        <Route path="/" element={<Layout />}>
          <Route path="/Home" element={<Home />} />
          <Route path="*" element={<NoPage />} />
        </Route>
      </Routes>
    </BrowserRouter>
  );
}

ReactDOM.render(
  <CookiesProvider>
    <App />
  </CookiesProvider>,
  document.getElementById("root")
);