import React, { useEffect, useMemo } from "react";
import useLocalStorage from "@rehooks/local-storage";
import debounce from "lodash.debounce";
import axios from "axios";

const Search = () => {
  const [accessToken] = useLocalStorage("accessToken", undefined);

  const handleOnchange = async (e) => {
    const params = new URLSearchParams([
      ["q", e.target.value],
      ["type", "track"],
    ]);

    const res = await axios.get("https://api.spotify.com/v1/search", {
      params,
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
        Authorization: `Bearer ${accessToken}`,
      },
    });
    console.log(res.data);
  };

  const debouncedChangeHandler = useMemo(
    () => debounce(handleOnchange, 300),
    // eslint-disable-next-line react-hooks/exhaustive-deps
    []
  );

  useEffect(() => {
    return () => {
      debouncedChangeHandler.cancel();
    };
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  if (!accessToken) {
    return <h1>Please login to see the search</h1>;
  }

  return (
    <div>
      <input type="text" onChange={debouncedChangeHandler} />
    </div>
  );
};

export default Search;
