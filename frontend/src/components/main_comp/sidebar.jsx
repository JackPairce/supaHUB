import { NavLink } from "react-router-dom";
import downloadsIcon from "../../icons/downloads.png";
import fileIcon from "../../icons/file-16.png";
import homeIcon from "../../icons/home.png";

function Sidebar() {
  return (
    <div className="sidebar">
      <NavLink
        to="/"
        className={({ isActive }) => (isActive ? "link active" : "link")}
      >
        <img src={homeIcon} alt="Home" />
        Home
      </NavLink>
      <NavLink
        to="/downloads"
        className={({ isActive }) => (isActive ? "link active" : "link")}
      >
        <img src={downloadsIcon} alt="Downloads" />
        Downloads
      </NavLink>
      <NavLink
        to="/file"
        className={({ isActive }) => (isActive ? "link active" : "link")}
      >
        <img src={fileIcon} alt="File" />
        File management
      </NavLink>
      <p>2024c, all rights are reserved</p>
    </div>
  );
}

export default Sidebar;
