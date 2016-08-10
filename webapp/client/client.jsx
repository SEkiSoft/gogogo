// Copyright (c) 2016 David Lu
// See License.txt

export default class Client {
  constructor () {
    this.url = '';
    this.defaultHeaders = {
      'X-Requested-With': 'XMLHttpRequest'
    };
  }

  handleResponse (successCallback, errorCallback, error, result) {
    if (error) {
      console.error(result.body.message); // eslint-disable-line no-console

      if (errorCallback) {
        errorCallback(error, result);
        return;
      }
    }

    if (successCallback) {
      successCallback(result);
    }
  }
}
