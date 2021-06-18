class UploadFile {
    constructor(fx){
        this.fx = fx;
        this.status = "pre_upload";
        this.result = "";
        this.time = Date.now();
    }

    async upload(password){
        this.status = "uploading";

        const formData = new FormData();
        formData.append("file", this.fx);

        // We pass the password as a GET variable here because if we want to parse a multipart request in Go,
        // we need to also hold on to the files that are part of that request. In theory this could fill up
        // the drive quickly or create situations where contraband is stored before password verification.
        let req = await fetch("/upload?password=" + encodeURIComponent(password), {
            method: 'POST',
            body: formData
        });

        this.result = await req.text();
        this.status = (! req.ok) ? "error" : "done";
    }
}

export default UploadFile;