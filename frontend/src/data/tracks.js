export const tracksReducerActions = { ADD_TRACKS: "ADD_TRACKS" };

export const tracksReducer = (state, action) => {
  switch (action.type) {
    case tracksReducerActions.ADD_TRACKS:
      return {
        ...state,
        ...action.payload,
      };
    default:
      return state;
  }
};
