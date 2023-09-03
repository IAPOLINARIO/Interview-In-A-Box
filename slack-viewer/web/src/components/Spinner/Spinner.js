import React from 'react';

const Spinner = () => {
    return (
        <div className="w-full flex items-center justify-center mt-32">
            <div
                className="inline-block h-8 w-8 animate-spin rounded-full border-4 border-solid border-fuchsia-950 border-current border-r-transparent align-[-0.125em] motion-reduce:animate-[spin_1.5s_linear_infinite]"
                role="status">
                  <span
                      className="!absolute !-m-px !h-px !w-px !overflow-hidden !whitespace-nowrap !border-0 !p-0 ![clip:rect(0,0,0,0)]"
                  >Loading...</span>
            </div>
        </div>
    );
};

export default Spinner;
