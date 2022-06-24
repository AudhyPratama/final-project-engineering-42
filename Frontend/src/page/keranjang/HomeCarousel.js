import React, { useState } from 'react'
// import 'bootstrap/dist/css/bootstrap.min.css';
import Carousel from 'react-bootstrap/Carousel';

const data = [
    {
        image: require('../../asset/img/Item/c1.jpg'),
        caption: "Caption",
        description: "Description Here"
    },
    {
        image: require('../../asset/img/Item/c2.jpg'),
        caption: "Caption",
        description: "Description Here"
    },
    {
        image: require('../../asset/img/Item/c3.jpg'),
        caption: "Caption",
        description: "Description Here"
    },
    {
        image: require('../../asset/img/Item/c4.jpg'),
        caption: "Caption",
        description: "Description Here"
    }
]


function HomeCarousel() {
    const [index, setIndex] = useState(0);
    const handleSelect = (selectedIndex, e) => {
        setIndex(selectedIndex);
    };

    return (
        <>
            {/* carousel */}
            <Carousel activeIndex={index} onSelect={handleSelect}>
                {data.map((slide, i) => {
                    return (
                        <Carousel.Item>
                            <img
                                className="d-block w-100"
                                src={slide.image}
                                alt="slider image"
                            />
                        </Carousel.Item>
                    )
                })}
            </Carousel>
            {/* carousel-end */}
            <br></br>            
        </>
    )
}

export default HomeCarousel