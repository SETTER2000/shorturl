import http from 'k6/http';
import {sleep} from 'k6';
import { URL } from 'https://jslib.k6.io/url/1.0.0/index.js';


// export const options = {
//   // Вместо того чтобы печатать --vus 10 и --duration 30с каждый раз,
//   // когда вы запускаете скрипт, вы можете включать параметры в свой js файл
//   // как здесь и теперь не надо запускать так: k6 run --vus 10 --duration 30s script.js
//   // а просто так: k6 run script.js
//   vus: 2000, // пользователей
//   duration: '30s', // длительность теста
// };
// https://k6.io/docs/using-k6/scenarios/
export const options = {
    summaryTrendStats: ['avg', 'min', 'med', 'max', 'p(95)', 'p(99)', 'p(99.99)', 'count'],
    scenarios: {
        example_scenario: {
            // name of the executor to use
            executor: 'shared-iterations',

            // common scenario configuration
            startTime: '10s',
            gracefulStop: '5s',
            env: {EXAMPLEVAR: 'testing'},
            tags: {example_tag: 'testing'},

            // executor-specific configuration
            vus: 10,
            iterations: 200,
            maxDuration: '10s',
        },
        // another_scenario: {
        //   /*...*/
        // },
    },
};
export default function () {
    // const domain = `${__ENV.BaseURL}/`
    // const res = http.get(url);
    // const url = new URL(domain);

    const url = new URL('https://k6.io');
    url.searchParams.append('utm_medium', 'organic');
    url.searchParams.append('utm_source', 'test');
    url.searchParams.append('multiple', ['foo', 'bar']);

    const res = http.get(url.toString());
    console.log(res)
    // console.log(url)
    // http.get(url);
    // http.get('http://test.k6.io');
    // http.get('http://localhost:8080');
    // console.log(randomItem([1, 2, 3, 4]));
    // sleep(randomIntBetween(1, 5)); // sleep between 1 and 5 seconds
    sleep(1);
}
