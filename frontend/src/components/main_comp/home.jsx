import { useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import pdfIcon from "../../icons/pdf_icon.png";
import pngIcon from "../../icons/png_icon.png";
import searchIcon from "../../icons/search.png";
import txtIcon from "../../icons/txt_icon.png";
import Sidebar from "./sidebar";

function Home() {
  const iconsTypes = {
    "application/pdf": pdfIcon,
    "text/plain": txtIcon,
    "image/png": pngIcon,
  };

  const array = [
    {
      name: "file1",
      id: 1,
      size: 1024,
      type: "application/pdf",
    },
    {
      name: "file2",
      id: 2,
      size: 1024,
      type: "text/plain",
    },
    {
      name: "file3",
      id: 3,
      size: 1024,
      type: "application/pdf",
    },
  ];

  const [files, setFiles] = useState([...array, ...array, ...array]);
  const [search, setSearch] = useState("");

  const searchChange = (event) => {
    const searchTerm = event.target.value;
    setSearch(searchTerm);
    const filteredFiles = array.filter((file) =>
      file.name.toLowerCase().includes(searchTerm.toLowerCase())
    );
    setFiles(filteredFiles);
  };

  localStorage.setItem("test", "jack");

  // let data = localStorage.getItem("test");
  // console.log(data);

  const navigate = useNavigate();
  // navigate("/register");

  return (
    <>
      <div className="content">
        <div className="input-box">
          <input type="text" onChange={searchChange} value={search} />
          <img src={searchIcon} alt="Search Icon" />
        </div>
        <div className="fileparent">
          {files.map((file) => (
            <div className="file" key={file.id}>
              <img
                src={iconsTypes[file.type] || searchIcon}
                alt={`${file.type} Icon`}
              />
              <div className="details">
                <p>Name: {file.name}</p>
                <p>Size: {file.size} bytes</p>
              </div>
              <Link>Download</Link>
            </div>
          ))}
        </div>
      </div>
      <Sidebar />
    </>
  );
}

export default Home;
