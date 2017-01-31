// Copyright 2017 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

func templates() map[string]string {
	return map[string]string{ 
        "app.yaml": "YXBwbGljYXRpb246IHt7LlJlbmRlcmVyLk5hbWV9fQp2ZXJzaW9uOiAxCnJ1bnRpbWU6IGdvCmFwaV92ZXJzaW9uOiBnbzEKaGFuZGxlcnM6Ci0gdXJsOiAvLioKICBzY3JpcHQ6IF9nb19hcHAKLSB1cmw6IC8KICBzdGF0aWNfZGlyOiBzdGF0aWM=",
        "client.go": "Ly8gR0VORVJBVEVEIEZJTEU6IERPIE5PVCBFRElUIQoKcGFja2FnZSB7ey5SZW5kZXJlci5QYWNrYWdlfX0KCmltcG9ydCAoCiAgImJ5dGVzIgogICJlbmNvZGluZy9qc29uIgogICJmbXQiCiAgIm5ldC9odHRwIgogICJzdHJpbmdzIgopCiAgCmNvbnN0IHNlcnZpY2UgPSAiaHR0cDovL2xvY2FsaG9zdDo4MDgwIgoKdHlwZSBDbGllbnQgc3RydWN0IHsKCVNlcnZpY2Ugc3RyaW5nCn0gCgpmdW5jIE5ld0NsaWVudChzZXJ2aWNlIHN0cmluZykgKkNsaWVudCB7CgljbGllbnQgOj0gJkNsaWVudHt9CgljbGllbnQuU2VydmljZSA9IHNlcnZpY2UKCXJldHVybiBjbGllbnQKfQoKLy8te3tyYW5nZSAuUmVuZGVyZXIuTWV0aG9kc319Ci8vLXt7aWYgZXEgLlJlc3VsdFR5cGVOYW1lICIifX0KZnVuYyAoY2xpZW50ICpDbGllbnQpIHt7LkNsaWVudE5hbWV9fSh7e3BhcmFtZXRlckxpc3QgLn19KSAoZXJyIGVycm9yKSB7Ci8vLXt7ZWxzZX19CmZ1bmMgKGNsaWVudCAqQ2xpZW50KSB7ey5DbGllbnROYW1lfX0oe3twYXJhbWV0ZXJMaXN0IC59fSkgKHJlc3VsdCAqe3suUmVzdWx0VHlwZU5hbWV9fSwgZXJyIGVycm9yKSB7Ci8vLXt7ZW5kfX0KCXBhdGggOj0gY2xpZW50LlNlcnZpY2UgKyAie3suUGF0aH19IgoJLy8te3tpZiBoYXNQYXJhbWV0ZXJzIC59fQoJLy8te3tyYW5nZSAuUGFyYW1ldGVyc1R5cGUuRmllbGRzfX0JCgkvLy17e2lmIGVxIC5Qb3NpdGlvbiAicGF0aCJ9fQoJcGF0aCA9IHN0cmluZ3MuUmVwbGFjZShwYXRoLCAieyIgKyAie3suSlNPTk5hbWV9fSIgKyAifSIsIGZtdC5TcHJpbnRmKCIldiIsIHt7LkpTT05OYW1lfX0pLCAxKQoJLy8te3tlbmR9fQoJLy8te3tlbmR9fQoJLy8te3tlbmR9fQoJLy8te3tpZiBlcSAuTWV0aG9kICJQT1NUIn19Cglib2R5IDo9IG5ldyhieXRlcy5CdWZmZXIpCglqc29uLk5ld0VuY29kZXIoYm9keSkuRW5jb2RlKHt7Ym9keVBhcmFtZXRlck5hbWUgLn19KQoJcmVxLCBlcnIgOj0gaHR0cC5OZXdSZXF1ZXN0KCJ7ey5NZXRob2R9fSIsIHBhdGgsIGJvZHkpCgkvLy17e2Vsc2V9fQoJcmVxLCBlcnIgOj0gaHR0cC5OZXdSZXF1ZXN0KCJ7ey5NZXRob2R9fSIsIHBhdGgsIG5pbCkKCS8vLXt7ZW5kfX0KCWlmIGVyciAhPSBuaWwgewoJCXJldHVybgoJfQoJcmVzcCwgZXJyIDo9IGh0dHAuRGVmYXVsdENsaWVudC5EbyhyZXEpCglpZiBlcnIgIT0gbmlsIHsKCQlyZXR1cm4gCgl9CglkZWZlciByZXNwLkJvZHkuQ2xvc2UoKQoJLy8te3tpZiBuZSAuUmVzdWx0VHlwZU5hbWUgIiJ9fQoJZGVjb2RlciA6PSBqc29uLk5ld0RlY29kZXIocmVzcC5Cb2R5KQoJcmVzdWx0ID0gJnt7LlJlc3VsdFR5cGVOYW1lfX17fQoJZGVjb2Rlci5EZWNvZGUocmVzdWx0KQoJLy8te3tlbmR9fQoJcmV0dXJuICAKfQoKLy8te3tlbmR9fQo=",
        "provider.go": "Ly8gR0VORVJBVEVEIEZJTEU6IERPIE5PVCBFRElUIQoKcGFja2FnZSB7ey5SZW5kZXJlci5QYWNrYWdlfX0KCnR5cGUgUHJvdmlkZXIgaW50ZXJmYWNlIHsKLy8te3tyYW5nZSAuUmVuZGVyZXIuTWV0aG9kc319Ci8vLXt7aWYgaGFzUGFyYW1ldGVycyAufX0KLy8te3tpZiBoYXNSZXNwb25zZXMgLn19CiAge3suUHJvY2Vzc29yTmFtZX19KHBhcmFtZXRlcnMgKnt7LlBhcmFtZXRlcnNUeXBlTmFtZX19LCByZXNwb25zZXMgKnt7LlJlc3BvbnNlc1R5cGVOYW1lfX0pIChlcnIgZXJyb3IpCi8vLXt7ZWxzZX19CiAge3suUHJvY2Vzc29yTmFtZX19KHBhcmFtZXRlcnMgKnt7LlBhcmFtZXRlcnNUeXBlTmFtZX19KSAoZXJyIGVycm9yKQovLy17e2VuZH19Ci8vLXt7ZWxzZX19Ci8vLXt7aWYgaGFzUmVzcG9uc2VzIC59fQogIHt7LlByb2Nlc3Nvck5hbWV9fShyZXNwb25zZXMgKnt7LlJlc3BvbnNlc1R5cGVOYW1lfX0pIChlcnIgZXJyb3IpCi8vLXt7ZWxzZX19CiAge3suUHJvY2Vzc29yTmFtZX19KCkgKGVyciBlcnJvcikKLy8te3tlbmR9fQovLy17e2VuZH19CQovLy17e2VuZH19Cn0K",
        "server.go": "Ly8gR0VORVJBVEVEIEZJTEU6IERPIE5PVCBFRElUIQoKcGFja2FnZSB7ey5SZW5kZXJlci5QYWNrYWdlfX0KCmltcG9ydCAoCgkiZW5jb2RpbmcvanNvbiIKCSJmbXQiCgkibG9nIgoJIm5ldC9odHRwIgoJInN0cmNvbnYiCgoJImdpdGh1Yi5jb20vZ29yaWxsYS9tdXgiCikKCmZ1bmMgaW50VmFsdWUocyBzdHJpbmcpICh2IGludDY0KSB7Cgl2LCBfID0gc3RyY29udi5QYXJzZUludChzLCAxMCwgNjQpCglyZXR1cm4gdgp9Cgp2YXIgcHJvdmlkZXIgUHJvdmlkZXIKCi8vIGhhbmRsZXJzCi8vLXt7cmFuZ2UgLlJlbmRlcmVyLk1ldGhvZHN9fQoKLy8ge3suRGVzY3JpcHRpb259fQpmdW5jIHt7LkhhbmRsZXJOYW1lfX0odyBodHRwLlJlc3BvbnNlV3JpdGVyLCByICpodHRwLlJlcXVlc3QpIHsKCWxvZy5QcmludGYoInt7Lk5hbWV9fSIpOwoJdmFyIGVyciBlcnJvcgoJLy8te3tpZiBoYXNQYXJhbWV0ZXJzIC59fQoJLy8gaW5zdGFudGlhdGUgdGhlIHBhcmFtZXRlcnMgc3RydWN0dXJlCgl2YXIgcGFyYW1ldGVycyB7ey5QYXJhbWV0ZXJzVHlwZU5hbWV9fQoJLy8te3tpZiBlcSAuTWV0aG9kICJQT1NUIn19CgkvLyBkZXNlcmlhbGl6ZSByZXF1ZXN0IGZyb20gcG9zdCBkYXRhCglkZWNvZGVyIDo9IGpzb24uTmV3RGVjb2RlcihyLkJvZHkpCgllcnIgPSBkZWNvZGVyLkRlY29kZSgmcGFyYW1ldGVycy57e2JvZHlQYXJhbWV0ZXJGaWVsZE5hbWUgLn19KQoJaWYgZXJyICE9IG5pbCB7CgkJZm10LkZwcmludGYodywgIkVSUk9SOiAldiIsIGVycikKCQlyZXR1cm4KCX0KCWxvZy5QcmludGYoIlBBUkFNRVRFUlMgJSt2IiwgcGFyYW1ldGVycykKCS8vLXt7ZW5kfX0KCS8vIGdldCByZXF1ZXN0IGZpZWxkcyBpbiBwYXRoIGFuZCBxdWVyeSBwYXJhbWV0ZXJzCgkvLy17e2lmIGhhc1BhdGhQYXJhbWV0ZXJzIC59fQoJdmFycyA6PSBtdXguVmFycyhyKQoJLy8te3tlbmR9fQoJLy8te3tpZiBoYXNGb3JtUGFyYW1ldGVycyAufX0KCXIuUGFyc2VGb3JtKCkKCS8vLXt7ZW5kfX0KCS8vLXt7cmFuZ2UgLlBhcmFtZXRlcnNUeXBlLkZpZWxkc319CQoJLy8te3tpZiBlcSAuUG9zaXRpb24gInBhdGgifX0KCWlmIHZhbHVlLCBvayA6PSB2YXJzWyJ7ey5KU09OTmFtZX19Il07IG9rIHsKCQlwYXJhbWV0ZXJzLnt7Lk5hbWV9fSA9IGludFZhbHVlKHZhbHVlKQoJfQoJLy8te3tlbmR9fQkKCS8vLXt7aWYgZXEgLlBvc2l0aW9uICJmb3JtZGF0YSJ9fQoJaWYgbGVuKHIuRm9ybVsie3suSlNPTk5hbWV9fSJdKSA+IDAgewoJCXBhcmFtZXRlcnMue3suTmFtZX19ID0gaW50VmFsdWUoci5Gb3JtWyJ7ey5KU09OTmFtZX19Il1bMF0pCgl9CgkvLy17e2VuZH19CgkvLy17e2VuZH19CgkvLy17e2VuZH19CgkvLy17e2lmIGhhc1Jlc3BvbnNlcyAufX0JCgkvLyBpbnN0YW50aWF0ZSB0aGUgcmVzcG9uc2VzIHN0cnVjdHVyZQoJdmFyIHJlc3BvbnNlcyB7ey5SZXNwb25zZXNUeXBlTmFtZX19CgkvLy17e2VuZH19CgkvLyBjYWxsIHRoZSBwcm9jZXNzb3IJCgkvLy17e2lmIGhhc1BhcmFtZXRlcnMgLn19CgkvLy17e2lmIGhhc1Jlc3BvbnNlcyAufX0KCWVyciA9IHByb3ZpZGVyLnt7LlByb2Nlc3Nvck5hbWV9fSgmcGFyYW1ldGVycywgJnJlc3BvbnNlcykKCS8vLXt7ZWxzZX19CgllcnIgPSBwcm92aWRlci57ey5Qcm9jZXNzb3JOYW1lfX0oJnBhcmFtZXRlcnMpCgkvLy17e2VuZH19CgkvLy17e2Vsc2V9fQoJLy8te3tpZiBoYXNSZXNwb25zZXMgLn19CgllcnIgPSBwcm92aWRlci57ey5Qcm9jZXNzb3JOYW1lfX0oJnJlc3BvbnNlcykKCS8vLXt7ZWxzZX19CgllcnIgPSBwcm92aWRlci57ey5Qcm9jZXNzb3JOYW1lfX0oKQoJLy8te3tlbmR9fQoJLy8te3tlbmR9fQkKCWlmIGVyciA9PSBuaWwgewoJLy8te3sgaWYgaGFzUmVzcG9uc2VzIC59fQoJCS8vLXt7IGlmIC5SZXNwb25zZXNUeXBlIHwgaGFzRmllbGROYW1lZE9LIH19CgkJLy8gd3JpdGUgdGhlIG5vcm1hbCByZXNwb25zZQkJCgkJZW5jb2RlciA6PSBqc29uLk5ld0VuY29kZXIodykKCQllbmNvZGVyLkVuY29kZShyZXNwb25zZXMuT0spCgkJLy8te3tlbmR9fQoJLy8te3tlbmR9fQoJfSBlbHNlIHsKCQkvLyB3cml0ZSBhbiBlcnJvciByZXNwb25zZS4uLgoJCWZtdC5GcHJpbnRmKHcsICJFUlJPUjogJXYiLCBlcnIpCgl9Cn0KLy8te3tlbmR9fQoKZnVuYyBJbml0aWFsaXplKHAgUHJvdmlkZXIpIHsKCXByb3ZpZGVyID0gcAoJdmFyIHJvdXRlciA9IG11eC5OZXdSb3V0ZXIoKXt7cmFuZ2UgLlJlbmRlcmVyLk1ldGhvZHN9fQoJcm91dGVyLkhhbmRsZUZ1bmMoInt7LlBhdGh9fSIsIHt7LkhhbmRsZXJOYW1lfX0pLk1ldGhvZHMoInt7Lk1ldGhvZH19Iil7e2VuZH19CiAgICBodHRwLkhhbmRsZSgiLyIsIHJvdXRlcikKfQoKZnVuYyBSdW4oKSB7CglodHRwLkxpc3RlbkFuZFNlcnZlKCI6ODA4MCIsIG5pbCkKfQo=",
        "types.go": "Ly8gR0VORVJBVEVEIEZJTEU6IERPIE5PVCBFRElUIQoKcGFja2FnZSB7ey5SZW5kZXJlci5QYWNrYWdlfX0KCi8vIHR5cGVzCi8vLXt7cmFuZ2UgLlJlbmRlcmVyLlR5cGVzfX0KdHlwZSB7ey5OYW1lfX0gc3RydWN0IHsge3tyYW5nZSAuRmllbGRzfX0KICB7ey5OYW1lfX0ge3suVHlwZX19e3tpZiBuZSAuSlNPTk5hbWUgIiJ9fSBganNvbjoie3suSlNPTk5hbWV9fSJge3tlbmR9fXt7ZW5kfX0KfQoKLy8te3tlbmR9fQ==",
    }
}