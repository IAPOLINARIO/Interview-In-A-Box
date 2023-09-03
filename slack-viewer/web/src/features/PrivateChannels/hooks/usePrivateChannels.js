import {config} from "../../../infrastructure/api/key.mappers";
import {useQuery} from "@tanstack/react-query";
import {API} from "../../../infrastructure/api/Api";
import {generatePath, dataTransformer} from "../../../infrastructure/utils/utils";
import PropTypes from "prop-types";

export const usePrivateChannels = ({searchByUser, from, to}) => {
    const { key, path } = config.channels.private();
    const generatedPath = generatePath(searchByUser, from, to, path);

    const result = useQuery([key, searchByUser, from, to], async () => await API.get({ path: generatedPath }), {
        enabled: searchByUser !== undefined && searchByUser !== '',
        suspense: true,
        refetchOnWindowFocus: false,
        refetchOnReconnect: false,
        retry: false
    })

    const { messages, groups } = dataTransformer(Array.isArray(result?.data) ? result.data.flatMap((item) => item.Mpims ?? []) : null);
    const errorMessage = result?.data && !Array.isArray(result?.data) ? result?.data.toString().split("Err:").slice(-1)[0] : null;

    return {
        privateChannels: messages ?? [],
        groups: groups ?? [],
        errorMessage: errorMessage,
        isLoading: result.isLoading,
        isFetching: result.isFetching
    }
}

usePrivateChannels.propTypes = {
    searchByUser: PropTypes.string,
    from: PropTypes.string,
    to: PropTypes.string
}
