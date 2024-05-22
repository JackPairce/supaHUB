import React, { useState } from 'react';
import MultiFiles from '../utility/multiFiles';

const Init = () => {
    const [isVisible, setIsVisible] = useState(true);

    const close = () => {
        setIsVisible(false); 
    };

    return (
        <>
            {isVisible && (
                <div className="initialise" >
                    <button className='closebtn' onClick={close}>X</button>
                    <h2>Please select the files that you want to initialise.</h2>
                    <MultiFiles />
                </div>
            )}
        </>
    );
};

export default Init;
