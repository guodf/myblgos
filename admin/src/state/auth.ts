import { AuthModel } from "../types";

class AuthState {
  authModel: AuthModel | null = null;
  login(authModel: AuthModel) {
    this.authModel = authModel;
    localStorage.setItem("blogs_auth", JSON.stringify(authModel));
  }
  logout() {
    localStorage.removeItem("blogs_auth");
    this.authModel = null;
  }

  isLogin(): boolean {
    return this.authModel != null
  }
}

const authState = new AuthState();

export default authState;
