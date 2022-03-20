import React, { useEffect } from "react";
import useLocalStorage from "@rehooks/local-storage";
import { Navigate, useLocation } from "react-router-dom";

const OAuthLoginCallback = () => {
  const [accessToken, setAccessToken] = useLocalStorage("accessToken", undefined);

  const { hash } = useLocation();
  const accessTokenParam = new URLSearchParams(hash.substring(1)).get("access_token");

  useEffect(() => {
    if (accessTokenParam) {
      setAccessToken(accessTokenParam);
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [accessTokenParam]);

  if (accessToken) {
    return <Navigate to="/" />;
  }

  return null;
};

export default OAuthLoginCallback;
