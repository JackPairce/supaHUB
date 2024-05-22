import { useEffect } from "react";
import { useNavigate } from "react-router-dom";

function Checker() {
  const navigate = useNavigate();

  const check = () => {
    const res = localStorage.getItem("userID");
    navigate(res ? "/home" : "/login");
  };

  useEffect(check, []);

  return <></>;
}

export default Checker;
