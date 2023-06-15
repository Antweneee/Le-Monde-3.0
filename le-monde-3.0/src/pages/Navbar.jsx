import React from "react";
import { useState, useEffect } from 'react';
import { useNavigate } from "react-router-dom";

function Navbar() {
    let navigate = useNavigate(); 
    const [hover, setHover] = useState(false)

    useEffect(() => {
      }, []);

    const hoverStyle = {
        cursor: "pointer",
        transition: "0.3s ease-in-out",
        transform: "scale(1.3)"
    }
    const normalStyle = {
        transition: "0.3s ease-in-out"
    }

    const onMouseEnter = () => {
        setHover(true)
    }
    const onMouseLeave = () => {
        setHover(false)
    }
    const onClick = () => {
        navigate("/");
    }

    return (
        <div class="">
            <nav class="text-center py-4 px-6 w-full mt-2"> 
                <i class="fa-solid fa-house fa-2xl text-black" onClick={onClick} style={hover ? hoverStyle : normalStyle} onMouseEnter={onMouseEnter} onMouseLeave={onMouseLeave}></i>
            </nav>
        </div>
    );
}

export default Navbar;

// Valentin Fouillet valouu.fouillet@gmail.com