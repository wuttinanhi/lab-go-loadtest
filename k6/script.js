import { check, sleep } from "k6";
import http from "k6/http";

export const options = {
  scenarios: {
    constant_request_rate: {
      executor: "constant-arrival-rate",
      rate: __ENV.COUNT, // number of iterations (requests) per timeUnit
      timeUnit: "1s", // iterations per second
      duration: __ENV.DURATION || "1s", // total duration of the test
      preAllocatedVUs: __ENV.COUNT || 1000, // to maintain the rate
      maxVUs: __ENV.COUNT || 1000, // maximum VUs needed
    },
  },
};

export default function () {
  const res = http.get(`${__ENV.URL}`);
  check(res, {
    "is status 200": (r) => r.status === 200,
  });
  sleep(1); // optional: if you want to add a delay between requests
}
