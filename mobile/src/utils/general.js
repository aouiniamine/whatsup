export function validateEmail($email) {
    var reg = /^([A-Za-z0-9_\-\.])+\@([A-Za-z0-9_\-\.])+\.([A-Za-z]{2,4})$/;
    return reg.test($email);
}
export function validateUsername($email) {
    var reg = /^([a-z0-9_\-\.]{8,20})$/;
    return reg.test($email);
}
