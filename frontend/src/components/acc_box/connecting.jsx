import Lottie from "lottie-react";
import { useEffect, useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import {
  TryConnect,
  Login as login_server,
} from "../../../wailsjs/go/main/App";
import animationData from "../../assets/hub-anim.json";

function Connecting(props) {
  const { type } = props;

  const navigate = useNavigate();

  const [url, setUrl] = useState({
    Address: "",
    Port: "8080",
  });
  const [isRadioChecked, setIsRadioChecked] = useState(false);
  const [wrongInput, setWrongInput] = useState("");
  const [connectionStatus, setConnectionStatus] = useState({
    msg: "",
    ok: false,
  });
  const [isloading, setIsloading] = useState(false);

  useEffect(() => {
    if (isRadioChecked)
      setUrl({
        ...url,
        Address: "localhost",
      });
    else
      setUrl({
        ...url,
        Address: "",
      });
  }, [isRadioChecked]);

  useEffect(() => {
    let add = localStorage.getItem("ServerUrl");
    if (add) {
      if (add == "localhost") setIsRadioChecked(true);
      else
        setUrl({
          ...url,
          Address: add,
        });
    }
  }, []);

  const handleTitle = (str) => {
    return str.charAt(0).toUpperCase() + str.slice(1);
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    if (
      type != "login" &&
      e.target.elements.password.value != e.target.elements.cpassword.value
    ) {
      setWrongInput("wrong password");
      return;
    }
    setIsloading(true);
    let pong = await TryConnect(
      e.target.elements.Address.value,
      e.target.elements.Port.value
    );
    setIsloading(false);
    if (pong) {
      setConnectionStatus({ msg: "Connected", ok: true });
      // alert("Connected");
      localStorage.setItem("ServerUrl", e.target.elements.Address.value);
      setIsloading(true);
      let id = await login_server(
        e.target.elements.nickname.value,
        e.target.elements.password.value
      );
      setIsloading(false);
      if (id != -1) {
        alert("Logged in");
        localStorage.setItem("userID", id);
        navigate("/home");
      } else {
        setWrongInput("Wrong Username or Password");
      }
    } else {
      setConnectionStatus({ msg: "Wrong credentials", ok: false });
    }
  };

  return (
    <>
      <Lottie animationData={animationData} className="logo" />
      <div className="Login" id="Login">
        <h2>{handleTitle(type)}</h2>
        <form onSubmit={handleSubmit}>
          <label for="nickname">Nickname:</label>
          <input type="text" name="nickname" required />

          <label for="password">Password:</label>
          <input type="password" name="password" required />

          {type != "login" ? (
            <>
              <label htmlFor="confPassword">Confirm your password:</label>
              <input type="password" name="cpassword" />
            </>
          ) : (
            <></>
          )}

          <p className="status error">{wrongInput}</p>

          <label htmlFor="Address">Address:</label>
          <div className="selector">
            <div>
              <input
                type="radio"
                id="network"
                name="drone"
                value="network"
                checked={!isRadioChecked}
                onClick={() => {
                  setIsRadioChecked(false);
                }}
              />
              <label for="network">Network</label>
            </div>

            <div>
              <input
                type="radio"
                id="local"
                name="drone"
                value="local"
                checked={isRadioChecked}
                onClick={(e) => {
                  setIsRadioChecked(true);
                }}
              />
              <label for="local">Local</label>
            </div>
          </div>

          <div className="ip">
            <input
              value={url.Address}
              onChange={(e) => {
                setUrl({
                  ...url,
                  Address: e.target.value,
                });
              }}
              className="Address"
              type="text"
              name="Address"
              required
              disabled={isRadioChecked}
            />
            <input
              value={url.Port}
              onChange={(e) => {
                setUrl({
                  ...url,
                  Port: e.target.value,
                });
              }}
              className="Port"
              type="text"
              name="Port"
              min="1"
              max="65535"
              required
            />
          </div>

          <p className={"status " + (connectionStatus.ok ? "good" : "error")}>
            {connectionStatus.msg}
          </p>

          {type == "login" ? (
            <p>
              New to SupaHub? <Link to="/register"> sign up</Link>
            </p>
          ) : (
            <p>
              Already have a SupaHub account? <Link to="/login">sign in</Link>
            </p>
          )}
          <button type="submit" data-status={!isloading ? "off" : ""}>
            {handleTitle(type)}
          </button>
        </form>
      </div>
    </>
  );
}

export default Connecting;
