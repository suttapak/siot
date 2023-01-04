import 'dart:convert';

import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:smiot/bloc/common_bloc.dart';
import 'package:smiot/models/boxes_models.dart';
import 'package:http/http.dart' as http;

class BoxBloc extends Bloc<MyEvent, MyState> {
  BoxBloc() : super(StateInitialized()) {
    on<GetBoxesEvent>(_onGetBoxes);
  }

  final _storage = const FlutterSecureStorage();

  Future<void> _onGetBoxes(GetBoxesEvent event, Emitter emit) async {
    emit(StateLoading());
    final token = await _storage.read(key: 'accessToken');
    final res = await http.get(
      Uri.parse('http://127.0.0.1:4000/boxes/members'),
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
    final boxes = json.map((e) => Boxes.fromJson(e)).toList();
    emit(GetBoxesStateSuccess(boxes));
  }
}

class GetBoxesEvent extends MyEvent {}

class GetBoxesStateSuccess extends MyState {
  final List<Boxes> boxes;

  GetBoxesStateSuccess(this.boxes);
}
