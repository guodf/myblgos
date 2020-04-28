export const formatTime = (seconds?: number) => {
  if (!seconds) {
    return "";
  }
  if (seconds?.toString().length === 10) {
    const date = new Date(seconds * 1000);
    return "y-m-d h:M:s".replace("y", date.getFullYear().toString())
      .replace("m", (date.getMonth() + 1).toString().padStart(2, '0'))
      .replace("d", date.getDate().toString().padStart(2, '0'))
      .replace("h", date.getHours().toString().padStart(2, '0'))
      .replace("M", date.getMinutes().toString().padStart(2, '0'))
      .replace("s", date.getSeconds().toString().padStart(2, '0'))
  }
}