import React, { useEffect, useMemo, useReducer } from "react";
import useLocalStorage from "@rehooks/local-storage";
import debounce from "lodash.debounce";
import axios from "axios";
import { errorHandler } from "../../utils/errorhandler";
import { tracksReducer, tracksReducerActions } from "../../data/tracks";

import styles from "./styles.module.css";

const Search = () => {
  const [accessToken] = useLocalStorage("accessToken", undefined);
  const [tracksState, dispatch] = useReducer(tracksReducer, {});

  const handleOnchange = async (e) => {
    const params = new URLSearchParams([
      ["q", e.target.value],
      ["type", "track"],
    ]);

    try {
      const res = await axios.get("https://api.spotify.com/v1/search", {
        params,
        headers: {
          Accept: "application/json",
          "Content-Type": "application/json",
          Authorization: `Bearer ${accessToken}`,
        },
      });

      dispatch({
        type: tracksReducerActions.ADD_TRACKS,
        payload: res.data.tracks,
      });
    } catch (error) {
      errorHandler(error);
    }
  };

  const debouncedChangeHandler = useMemo(
    () => debounce(handleOnchange, 300),
    // eslint-disable-next-line react-hooks/exhaustive-deps
    []
  );

  const tracksList = useMemo(() => tracksState.items || [], [tracksState.items]);

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

      <ul className={styles.list}>
        {tracksList.map((track) => (
          <li key={track.id} className={styles.listItem}>
            <img className={styles.cover} src={track.album.images[0].url} alt={track.name} />
            <p className={styles.name}>{track.name}</p>
            <audio controls>
              <source src={track.preview_url} type="audio/mpeg" />
              Your browser does not support the audio tag.
            </audio>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default Search;
