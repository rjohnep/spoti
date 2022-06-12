export const errorHandler = (error) => {
  if (error?.response?.data?.error?.message === "The access token expired") {
    window.location.href = "/login";
  }
};
