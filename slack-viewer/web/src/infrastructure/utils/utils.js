export const accountingFormatNumber = (
    value,
    locale = 'en-US'
) => {
    const formatter = new Intl.NumberFormat(locale, {
        style: 'decimal',
        currency: undefined,
        currencySign: undefined
    })

    return formatter.format(value)
}

export const generatePath = (searchByUser, from ,to, path) => {
        let queryParams = '';
        let date, month, year;

        if (searchByUser !== null && searchByUser !== undefined && searchByUser.length) {
            queryParams += `username=${searchByUser}`
        }
        if (from) {
            queryParams += `&startdate=${formatDate(from)}`
        }
        if (to) {
            queryParams += `&enddate=${formatDate(to)}`
        }
        return path + (queryParams !== '' ? `?${queryParams}` : '')
}

const formatDate = (inputDate)  => {

  const date = new Date(inputDate);
  const month = (date.getMonth() + 1).toString().padStart(2, '0');
  const year = date.getFullYear();

  return `${month}${year}`;
}

export const highlightPortionOfTheText = (message, criteria) => {
    const escapedSearch = criteria.replace(/[.*+?^${}()|[\]\\]/g, '\\$&');
    const regex = new RegExp(escapedSearch, 'gi');
    return message.replace(regex, '<span class="bg-amber-200">$&</span>');
}

export const truncate = (text, cuttingLength) =>
text.length > 40 ? `${text.substring(0, cuttingLength ?? 15)}...` : text



export const dataTransformer = (data) => {

    if (data === null) return { messages: [], groups: [] };

    const flattenedMessages = data
      .sort((a, b) => {
        if (a.GroupName < b.GroupName) {
          return -1;
        } else if (a.GroupName > b.GroupName) {
          return 1;
        } else {
          return 0;
        }
      })
      .flatMap((item) =>
      item.Groups.flatMap((group) => {
          return group.Messages.map((message) => ({
            ...message,
            groupName: group.Context
          }));
        }
      )
    );

    let messages = flattenedMessages.map((message) => {
      const timestamp = parseInt(message.ts.split('.')[0]) * 1000;
      const humanReadableDate = new Date(timestamp).toLocaleString();
      const fullName = message.user_profile.real_name;
      const displayName = message.user_profile.display_name;
      const avatar = message.user_profile.image_72;
      return {
        key: message.ts,
        message: message.text,
        timestamp,
        userId: message.user,
        humanReadableDate,
        groupName: message.groupName,
        user: {
          fullName,
          displayName,
          avatar,
        },
      };
    });
    
    const groups = Array.from(
      new Set(messages.map((data) => data.groupName)),
    ).sort((a, b) => {
      if (a < b) {
        return -1;
      } else if (a > b) {
        return 1;
      } else {
        return 0;
      }
    });

    return { messages, groups };
  };