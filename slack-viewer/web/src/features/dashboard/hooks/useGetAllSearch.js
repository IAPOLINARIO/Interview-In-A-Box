import {config} from "../../../infrastructure/api/key.mappers";
import {useQuery} from "@tanstack/react-query";
import {API} from "../../../infrastructure/api/Api";
import PropTypes from "prop-types";
import {generatePath, dataTransformer} from "../../../infrastructure/utils/utils";

export const useGetAllSearch = ({searchByUser, from, to}) => {
    const {key, path} = config.allSearch.default();
    const generatedPath = generatePath(searchByUser, from, to, path);

    const result = useQuery([key, searchByUser, from, to], async () => await API.get({path: generatedPath}), {
        enabled: searchByUser !== undefined && searchByUser !== '',
        suspense: true,
        refetchOnWindowFocus: false,
        refetchOnReconnect: false,
        retry: false
    })

    let data = null;

    if (Array.isArray(result?.data)) {

        const publicData = result.data.flatMap((item) => item.Channels ?? []);
        const privateData = result.data.flatMap((item) => item.Mpims ?? []);
        const directMessagesData = result.data.flatMap((item) => item.DMs ?? []);
        const groups = result.data.flatMap((item) => item.Groups ?? []);

        data = [
            ...Object.values(publicData),
            ...Object.values(privateData),
            ...Object.values(directMessagesData),
            ...Object.values(groups),
          ];
    }
    
    const { messages, groups } = dataTransformer(data);
    const errorMessage = result?.data && !Array.isArray(result?.data) ? result?.data.toString().split("Err:").slice(-1)[0] : null;

    return {
        allData: messages ?? [],
        groups: groups ?? [],
        errorMessage: errorMessage,
        isLoading: result?.isLoading ?? false,
        isFetching: result?.isFetching ?? false
    }
}

useGetAllSearch.propTypes = {
    searchByUser: PropTypes.string,
    from: PropTypes.string,
    to: PropTypes.string
}
