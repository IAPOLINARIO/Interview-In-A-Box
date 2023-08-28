# GoLang Log Parser Challenge - Senior Software Engineer Position

## Objective

Develop a concurrent log parsing application in Go that can efficiently process large log files. The application should identify specific patterns in the log lines, categorize them, and provide basic analytics.

## Log File Format

Logs will be in the following format:

```
[YYYY-MM-DD HH:MM:SS] [SEVERITY] [MODULE]: Message details...
```

Example:

```
[2023-08-04 12:30:45] [INFO] [AUTH]: User John logged in successfully.
[2023-08-04 12:31:05] [ERROR] [API]: Connection timeout for request /getDetails.
[2023-08-04 12:32:30] [DEBUG] [DB]: Executing query SELECT * FROM users.

```

## Requirements

1. Log Level Filtering:
   The application should allow the user to specify which log levels (e.g., INFO, ERROR, DEBUG) they're interested in.

2. Module-specific Parsing:
   Allow the user to filter results based on specific modules (e.g., AUTH, API, DB).

3. Time Range Filtering:
   The user should be able to specify a start and end time, and the application should return logs only from that time range.

4. Analytics:
   Provide analytics like:

- Number of logs per severity level.
- Number of logs per module.
- Peak time with the highest log entries.

5. Concurrency:
   Make use of goroutines and channels to process the logs concurrently, ensuring efficient use of resources.

6. Output Format:
   Allow the user to specify if they want the output displayed in the console or written to an output file.

## Scenarios to Handle

1. Large Files:
   The application should be capable of processing log files that are several gigabytes in size without crashing or consuming all available memory.

2. Invalid Log Format:
   Handle logs that don't match the specified format gracefully. These should be recorded as 'malformed' logs.

3. User Input Validation:
   Ensure that user inputs for filtering (like date, time, log level) are validated properly before processing.

## Bonus

- Implement a user-friendly CLI interface for the tool.
- Provide unit tests for your application.
- Implement a mechanism that watches a log file in real-time and processes new log entries as they are written.

## Sample Logs for Testing

```
[2023-08-04 12:30:45] [INFO] [AUTH]: User Alice logged in successfully.
[2023-08-04 12:30:50] [DEBUG] [AUTH]: User Bob password validation failed.
[2023-08-04 12:31:05] [ERROR] [API]: Connection timeout for request /getData.
[2023-08-04 12:31:10] [INFO] [API]: Request /updateProfile successful for User Alice.
[2023-08-04 12:32:30] [DEBUG] [DB]: Executing query SELECT name FROM products.
[2023-08-04 12:32:35] [ERROR] [DB]: Connection lost to the database.

```
