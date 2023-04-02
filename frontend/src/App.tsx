import Conversation from "./Conversation";
import SideBar from "./SideBar";
import "./styles/scrollbar.css";
import Box from "@mui/material/Box";
import CssBaseline from "@mui/material/CssBaseline";
import * as React from "react";
import { ToastContainer } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";

export default function App() {
  return (
    <Box sx={{ display: "flex" }}>
      <CssBaseline />
      <ToastContainer />
      <SideBar />
      <Conversation />
    </Box>
  );
}
