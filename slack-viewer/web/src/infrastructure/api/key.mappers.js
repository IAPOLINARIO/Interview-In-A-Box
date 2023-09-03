export const config = {
    dms: {
        all: () => ['dms'],
        getAllDms: () => ({
            key: [...config.dms.all(), 'all'],
            path: '/slack/user/report'
        }),
    },
    channels: {
        all: () => ['channels'],
        private: () => ({
            key: [...config.channels.all(), 'all/private'],
            path: '/slack/user/report'
        }),
        public: () => ({
            key: [...config.channels.all(), 'all/public'],
            path: '/slack/user/report'
        })
    },
    groups: {
        all: () => ['groups'],
        dms: () => ({
            key: [...config.groups.all(), 'dms'],
            path: '/slack/user/report'
        })
    },
    allSearch: {
        default: () => ({
            key: ['search-all', 'default'],
            path: '/slack/user/report'
        })
    }
}
