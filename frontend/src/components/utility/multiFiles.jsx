import React, { useState, useRef } from 'react';
import pngIcon from '../../icons/png_icon.png';
import pdfIcon from '../../icons/pdf_icon.png';
import txtIcon from '../../icons/txt_icon.png';


const MultiFiles = () => {

  const iconsTypes = {
    "application/pdf": pdfIcon,
    "text/plain": txtIcon,
    "image/png": pngIcon
  }

  const [files, setFiles] = useState([]);
  const [showFiles, setShowFiles] = useState(false);
  const inputRef = useRef();

  const handleChange = (event) => {
    setFiles(prevFiles => [...prevFiles, ...Array.from(event.target.files)]);
  };

  const handleExpose = () => {
    setShowFiles(true);
    inputRef.current.click();
  };

  const deleteFile = (index) => {
    setFiles(prevFiles => prevFiles.filter((file, i) => i !== index));
  };

  return (
    <div className='choose'>
      <div className='file-box'>
        <input type="file" name='file' ref={inputRef} webkitdirectory mozdirectory multiple onChange={handleChange} style={{ display: 'none' }} />
        <button onClick={handleExpose}>Upload files</button>
      </div>
      {showFiles && files.length > 0 && (
        <div className='fileparent'>
          {files.map((file, index) => (
            <div className='file' key={index} >
              <button className='delete' onClick={() => deleteFile(index)}>X</button>
              <img
                src={iconsTypes[file.type]}
                alt={`${file.type} Icon`}
              />
              <p>{file.name}</p>
            </div>
          ))}
        </div>
      )}
    </div>
  );
};

export default MultiFiles;
