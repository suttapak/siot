import 'dart:convert';

import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:smiot/bloc/common_bloc.dart';
import 'package:http/http.dart' as http;

class AuthBloc extends Bloc<MyEvent, MyState> {
  static const _accessTokenKey = 'accessToken';
  final _storage = const FlutterSecureStorage();

  AuthBloc() : super(StateInitialized()) {
    on<LoginEvet>((event, emit) async =>
        _onLogin(event, emit, email: event.email, password: event.password));
    on<GetTokenEvent>(_onGetToken);
    on<LogoutEvent>(_onLogout);
  }
  void _onLogin(LoginEvet evet, Emitter emit,
      {required String email, required String password}) async {
    emit(StateLoading());

    var res = await http.post(Uri.parse('http://127.0.0.1:4000/auth/login'),
        body: jsonEncode({'email': email, 'password': password}));

    if (res.statusCode != 200) {
      emit(StateError(message: res.statusCode.toString()));
      return;
    }

    final json = jsonDecode(res.body) as Map<String, dynamic>;
    final accessToken = json['accessToken'].toString();
    await _storage.write(key: _accessTokenKey, value: accessToken.toString());
    emit(LoginStateSuccess(accessToken));
  }

  void _onGetToken(GetTokenEvent event, Emitter emit) async {
    emit(StateLoading());
    try {
      final token = await _storage.read(key: _accessTokenKey);
      if (token == null || token.isEmpty) {
        emit(StateError(message: 'unexist token'));
        return;
      }
      emit(GetTokenStateSubccess(accessToken: token));
    } catch (e) {
      emit(StateError(message: e.toString()));
    }
  }

  void _onLogout(LogoutEvent event, Emitter emit) async {
    emit(StateLoading());
    try {
      await _storage.delete(key: _accessTokenKey);

      emit(LogoutStateSuccess());
    } catch (e) {
      emit(StateError(message: e.toString()));
    }
  }
}

class AuthState extends MyState {
  final String token;

  AuthState(this.token);
}

class LoginEvet extends MyEvent {
  final String email;
  final String password;

  LoginEvet({required this.email, required this.password});
}

class LogoutEvent extends MyEvent {}

class LogoutStateSuccess extends AuthState {
  LogoutStateSuccess() : super('');
}

class GetTokenEvent extends MyEvent {}

class GetTokenStateSubccess extends AuthState {
  final String accessToken;

  GetTokenStateSubccess({required this.accessToken}) : super(accessToken);
}

class LoginStateSuccess extends AuthState {
  final String accessToken;

  LoginStateSuccess(this.accessToken) : super(accessToken);
}
