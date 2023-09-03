import {sizeFeatureSelector} from "./button.utils";

const PrimaryContainedButton = ({
                                    size = 'base',
                                    children,
                                    fullWidth = false,
                                    ...rest
                                }) => {
    const buttonSize = sizeFeatureSelector[size]
    return (
        <button
            type="button"
            {...rest}
            className={`${buttonSize} 
       text-white inline-block rounded-lg bg-gradient-to-br hover:bg-gradient-to-bl focus:ring-2 focus:outline-none 
       focus:ring-fuchsia-700 bg-fuchsia-800 hover:bg-fuchsia-900 disabled:opacity-75 font-semibold ${
                Boolean(fullWidth) ? 'w-full' : ''
            }`}
        >
            {children}
        </button>
    )
}

export default PrimaryContainedButton
