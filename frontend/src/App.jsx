import { BrowserRouter, Route, Routes } from "react-router-dom";
// import { Greet } from "../wailsjs/go/main/App";
import Connecting from "./components/acc_box/connecting";
// import Register from "./components/acc_box/register";
import Checker from "./components/main_comp/checker";
import Downloads from "./components/main_comp/download";
import Home from "./components/main_comp/home";
import Files from "./components/utility/files";
import "./sass/app.scss";

function App() {
  // Greet().then((res) => {
  //   console.log(res);
  // });
  // go to homewindow.location.href = "/home";

  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Checker />} exact />
        <Route path="/home" element={<Home />} exact />
        <Route path="/downloads" element={<Downloads />} exact />
        <Route path="/file" element={<Files />} exact />
        <Route path="/login" element={<Connecting type="login" />} exact />
        <Route
          path="/register"
          element={<Connecting type="register" />}
          exact
        />
        {/* <Route path="/register" element={<Register />} exact /> */}
      </Routes>
    </BrowserRouter>
  );
}

export default App;
