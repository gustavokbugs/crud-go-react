import axios from "axios";

const axiosService = axios.create({
  baseURL: "http://localhost:3333",
  })

export default axiosService