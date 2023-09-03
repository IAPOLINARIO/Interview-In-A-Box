import {Navigate, useOutlet} from "react-router-dom";
import {useOktaAuth} from "@okta/okta-react";

const PublicMainLayout = () => {

    const { authState } = useOktaAuth();
    const outlet = useOutlet()

    if (authState && authState.isAuthenticated) {
        return <Navigate to="/admin"/>
    }

    return (
        <div className="relative">
            <main>
                {outlet}
            </main>
        </div>
    )
}

export default PublicMainLayout