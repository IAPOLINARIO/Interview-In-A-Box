import React from 'react';
import PropTypes from "prop-types";

const TitleWithCountHeader = ({ title, count }) => {
    return (
        <div className="py-4 px-2 bg-white mb-3 relative">
            <p className="text-xl flex items-center font-bold before:content-[' '] before:h-14 before:w-1 before:absolute before:left-0 pl-3 before:bg-fuchsia-800">
                <span>{title}</span>&nbsp;-&nbsp;
                <span className="text-fuchsia-800">{count} records</span>
            </p>
        </div>
    );
};


TitleWithCountHeader.propTypes = {
    title: PropTypes.string.isRequired,
    count: PropTypes.number.isRequired
}

export default TitleWithCountHeader;
