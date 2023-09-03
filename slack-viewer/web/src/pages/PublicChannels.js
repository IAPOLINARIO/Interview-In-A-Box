import React, {lazy, Suspense} from 'react';
import PrimaryContainedButton from "../components/Buttons/PrimaryContainedButton";
import Spinner from "../components/Spinner/Spinner";
import {ErrorBoundary} from "react-error-boundary";
import {useQueryErrorResetBoundary} from "@tanstack/react-query";

const PublicChannelsContent = lazy(() => import("../features/PublicChannels/PublicChannelsContent"));

const PublicChannels = () => {
    const {reset} = useQueryErrorResetBoundary()

    return (
        <ErrorBoundary
            onReset={reset}
            fallbackRender={({resetErrorBoundary, error}) => (
                <div
                    role="alert"
                    className="flex flex-col justify-center items-center"
                >
                    <p className="text-center text-xl font-bold">
                        Something went wrong:
                    </p>
                    <pre className="text-center text-rose-500 text-lg">
                {error.message}
              </pre>
                    <PrimaryContainedButton onClick={resetErrorBoundary}>
                        Try again
                    </PrimaryContainedButton>
                </div>
            )}
        >
            <Suspense fallback={<Spinner/>}>
                <PublicChannelsContent />
            </Suspense>
        </ErrorBoundary>
    );
};

export default PublicChannels;
