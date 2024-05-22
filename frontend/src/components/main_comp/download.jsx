import { useState } from "react";
import pdfIcon from "../../icons/pdf_icon.png";
import pngIcon from "../../icons/png_icon.png";
import txtIcon from "../../icons/txt_icon.png";
import Sidebar from "./sidebar";

const Downloads = () => {
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
  return (
    <>
      <div className="download_content">
        {files.map((file) => (
          <div className="file">
            <div>
              <img src={iconsTypes[file.type]} alt={`${file.type} Icon`} />
              <p>{file.name}</p>
            </div>
            <progress value="20" min="0" max="100"></progress>
            <p>{file.size} bytes</p>
          </div>
        ))}
      </div>
      <Sidebar />
    </>
  );
};
export default Downloads;
