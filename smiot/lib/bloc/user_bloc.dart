import 'dart:convert';

import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:smiot/bloc/common_bloc.dart';
import 'package:http/http.dart' as http;
import 'package:smiot/models/user_models.dart';

class UserBloc extends Bloc<MyEvent, MyState> {
  UserBloc() : super(StateInitialized()) {
    on<GetUserEvent>(_onGetUser);
  }
  final _storage = const FlutterSecureStorage();

  Future<void> _onGetUser(GetUserEvent event, Emitter emit) async {
    emit(StateLoading());

    final token = await _storage.read(key: 'accessToken');
    final res = await http.get(
        Uri.parse('http://localhost:4000/user/${event.userId}'),
        headers: <String, String>{
          'context-type': 'application/json',
          'Authorization': 'Bearer $token'
        });

    if (res.statusCode != 200) {
      emit(StateError(message: res.body));
      return;
    }
    final json = jsonDecode(res.body) as Map<String, dynamic>;
    final user = Users.fromJson(json);
    emit(GetUserStateSuccess(user));
  }
}

class GetUserEvent extends MyEvent {
  final String userId;

  GetUserEvent({required this.userId});
}

class GetUserStateSuccess extends MyState {
  final Users user;

  GetUserStateSuccess(this.user);
}
