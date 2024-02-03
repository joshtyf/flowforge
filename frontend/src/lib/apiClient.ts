import axios from "axios"

// TODO: Insert prod env URL
const baseURL =
  process.env.NEXT_PUBLIC_APP_ENV === "DEV" ? "http://localhost:8080" : ""
const apiClient = axios.create({
  baseURL: `${baseURL}/api`,
  headers: {
    "Content-type": "application/json",
  },
})

export default apiClient
