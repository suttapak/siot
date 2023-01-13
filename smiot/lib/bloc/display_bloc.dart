import 'dart:convert';

import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:smiot/bloc/common_bloc.dart';
import 'package:smiot/models/display_models.dart';
import 'package:http/http.dart' as http;

class DisplayBloc extends Bloc<MyEvent, MyState> {
  DisplayBloc() : super(StateInitialized()) {
    on<GetDisplayEvent>(_onGetDisplays);
  }

  final _storage = const FlutterSecureStorage();

  Future<void> _onGetDisplays(GetDisplayEvent event, Emitter emit) async {
    try {
      emit(StateLoading());
      final token = await _storage.read(key: 'accessToken');
      final res = await http.get(
        Uri.parse(
            'https://api.rocket-translate.com/boxes/${event.boxId}/displays'),
        headers: <String, String>{
          'context-type': 'application/json',
          'Authorization': 'Bearer $token'
        },
      );
      if (res.statusCode != 200) {
        emit(StateError(message: res.body));
        return;
      }
      final json = jsonDecode(res.body) as List;
      final displays = json.map((e) => Displays.fromJson(e)).toList();
      emit(GetDisplaysStateSuccess(displays));
    } catch (e) {
      emit(StateError(message: 'get display error'));
    }
  }
}

class GetDisplayEvent extends MyEvent {
  final String boxId;

  GetDisplayEvent(this.boxId);
}

class GetDisplaysStateSuccess extends MyState {
  final List<Displays> displays;

  GetDisplaysStateSuccess(this.displays);
}
