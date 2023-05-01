import axios from 'axios';

/**
 * 以 Json 方式发送 POST 请求
 * @param {String} url 请求的 Url
 * @param {Object} data 请求时携带的数据
 */

// const headerJSON = { "Content-Type": "application/json" }
//
export function putByJson(url: string, data: object) {
	if (import.meta.env.DEV) {
		url = 'http://localhost:9000' + url;
	}
	return new Promise<any>((resolve, reject) => {
		console.log('putByJson: ', JSON.stringify(data));
		axios
			.put(url, JSON.stringify(data), {
				headers: {
					Authorization: 'Bearer ' + localStorage.getItem('prj-jwt'),
					'Content-Type': 'application/json'
				}
			})
			.then((res) => {
				resolve(res);
			})
			.catch((err) => {
				reject(err);
			});
	});
}


export function postByJson(url: string, data: object) {
	if (import.meta.env.DEV) {
		url = 'http://localhost:9000' + url;
	}
	return new Promise<any>((resolve, reject) => {
		console.log('postByJson: ', JSON.stringify(data));
		axios
			.post(url, JSON.stringify(data), {
				headers: {
					Authorization: 'Bearer ' + localStorage.getItem('prj-jwt'),
					'Content-Type': 'application/json'
				}
			})
			.then((res) => {
				resolve(res);
			})
			.catch((err) => {
				reject(err);
			});
	});
}

export function get(url: string) {
	if (import.meta.env.DEV) {
		url = 'http://localhost:9000' + url;
	}
	return new Promise<any>((resolve, reject) => {
		axios
			.get(url, { headers: { Authorization: 'Bearer ' + localStorage.getItem('prj-jwt') } })
			.then((res) => {
				resolve(res);
			})
			.catch((err) => {
				reject(err);
			});
	});
}