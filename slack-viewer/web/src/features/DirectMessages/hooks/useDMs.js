import {config} from "../../../infrastructure/api/key.mappers";
import {useQuery} from "@tanstack/react-query";
import {API} from "../../../infrastructure/api/Api";
import PropTypes from "prop-types";
import {generatePath, dataTransformer} from "../../../infrastructure/utils/utils";

export const useGetDMs = ({searchByUser, from, to}) => {
    const {key, path} = config.dms.getAllDms();
    const generatedPath = generatePath(searchByUser, from, to, path);

    const result = useQuery([key, searchByUser, from, to], async () => await API.get({path: generatedPath}), {
        enabled: searchByUser !== undefined && searchByUser !== '',
        suspense: true,
        refetchOnWindowFocus: false,
        refetchOnReconnect: false,
        retry: false
    })

    const { messages, groups } = dataTransformer(Array.isArray(result?.data) ? result.data.flatMap((item) => item.DMs ?? []) : null);
    const errorMessage = result?.data && !Array.isArray(result?.data) ? result?.data.toString().split("Err:").slice(-1)[0] : null;

    return {
        dms: messages ?? [],
        groups: groups ?? [],
        errorMessage: errorMessage,
        isLoading: result.isLoading,
        isFetching: result.isFetching
    }
}

useGetDMs.propTypes = {
    searchByUser: PropTypes.string,
    from: PropTypes.string,
    to: PropTypes.string
}
